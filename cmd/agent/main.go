package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	baseImage := "ubuntu:22.04"
	err = pullImage(cli, baseImage)
	if err != nil {
		panic(err)
	}

	script, err := ioutil.ReadFile("./scripts/angular_build.sh")
	if err != nil {
		panic(err)
	}

	err = runScriptOnImage(cli, baseImage, string(script))
	if err != nil {
		panic(err)
	}
}

func pullImage(cli *client.Client, imageName string) error {
	fmt.Printf("Pulling %s image...\n", imageName)

	// Pull the image
	reader, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()

	// Read and print the pull output
	output, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	fmt.Println(string(output))

	fmt.Printf("%s image pulled successfully.\n", imageName)
	return nil
}

func runScriptOnImage(cli *client.Client, imageName string, script string) error {
	fmt.Printf("Running bash script on %s image...\n", imageName)

	// Create a container with the specified image
	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: imageName,
			Cmd:   []string{"bash", "-c", script},
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
	out, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{ShowStdout: true})
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

	fmt.Printf("Bash script executed successfully on %s image.\n", imageName)
	return nil
}
