package pipeline

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
)

func RunPipeline(definition *Definition) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Create a container with the specified image
	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: definition.Image,
			Cmd:   []string{"bash", "-c", definition.GetScript()},
		},
		nil,
		nil,
		nil,
		"",
	)
	if err != nil {
		return err
	}

	// Start the container
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	// Wait for the container to finish
	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	// Retrieve the container logs
	out, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})

	if err != nil {
		return err
	}
	defer out.Close()

	// Print the container logs
	logs, err := io.ReadAll(out)
	if err != nil {
		return err
	}
	fmt.Println(string(logs))

	return nil
}
