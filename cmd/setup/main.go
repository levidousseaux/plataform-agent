package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"io"
	"os"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	err = buildDockerImage(cli, "./scripts/spa/", "spa_build_env")
	if err != nil {
		panic(err)
	}
}

func buildDockerImage(cli *client.Client, buildContextPath, imageName string) error {
	buildContext, err := archive.Tar(buildContextPath, archive.Uncompressed)
	if err != nil {
		return err
	}

	buildResponse, err := cli.ImageBuild(
		context.Background(),
		buildContext,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{imageName},
		},
	)
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()

	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}
