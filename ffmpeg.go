package goradiru

import (
	"os"
	"os/exec"
)

type ffmpeg struct {
	*exec.Cmd
}

func newFFMPEG(inputFilePath string) (*ffmpeg, error) {
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	/* #nosec */
	cmd := exec.Command(cmdPath, "-i", inputFilePath)
	cmd.Env = os.Environ()
	return &ffmpeg{cmd}, nil
}

func (f *ffmpeg) setArgs(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *ffmpeg) execute(output string) ([]byte, error) {
	f.Args = append(f.Args, output)
	return f.CombinedOutput()
}
