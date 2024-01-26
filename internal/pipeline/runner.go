package pipeline

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"os"
	"strings"
)

func RunPipeline(definition *Definition) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	id, err := StartContainer(ctx, cli, definition.Image)
	if err != nil {
		return err
	}

	err = RunStages(ctx, cli, id, definition)
	if err != nil {
		fmt.Println(err)
	}

	err = cli.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}

func StartContainer(ctx context.Context, cli *client.Client, image string) (string, error) {
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{Image: image, Tty: true},
		nil,
		nil,
		nil,
		"",
	)
	if err != nil {
		return "", err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func RunStages(ctx context.Context, cli *client.Client, containerId string, definition *Definition) error {
	for _, stage := range definition.Stages {
		err := RunStage(ctx, cli, containerId, stage)

		if err != nil {
			return err
		}
	}

	return nil
}

func RunStage(ctx context.Context, cli *client.Client, containerId string, stage Stage) error {
	fmt.Printf("[%s]\n", stage.Name)

	for _, step := range stage.Steps {
		if err := RunStep(ctx, cli, containerId, step); err != nil {
			return err
		}
	}

	fmt.Println("")

	return nil
}

func RunStep(context context.Context, cli *client.Client, containerId string, step Step) error {
	fmt.Printf("--Step %s\n", step.Name)
	script := strings.Join(step.Commands, "\n")

	fmt.Println("----Running: ", script)

	execCreateResp, err := cli.ContainerExecCreate(context, containerId, types.ExecConfig{
		Cmd:          []string{"bash", "-c", script},
		AttachStdout: true,
		AttachStderr: true,
	})

	if err != nil {
		return err
	}

	attachResp, err := cli.ContainerExecAttach(context, execCreateResp.ID, types.ExecStartCheck{
		Detach: false,
	})

	if err != nil {
		return err
	}
	defer attachResp.Close()

	stdoutChan := make(chan string)

	go func() {
		var stdoutBuf bytes.Buffer
		_, err := io.Copy(io.MultiWriter(os.Stdout, &stdoutBuf), attachResp.Reader)
		if err != nil {
			panic(err)
		}
		stdoutChan <- stdoutBuf.String()
	}()

	waitResp, err := cli.ContainerExecInspect(context, execCreateResp.ID)
	if err != nil {
		return err
	}

	<-stdoutChan

	// TODO: npm scripts is not returning exit code 1
	if waitResp.ExitCode != 0 {
		return errors.New("exit code is not 0")
	}

	return nil
}
