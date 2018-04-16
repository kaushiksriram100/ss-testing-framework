//package models
//Author: Sriram Kaushik
//Date: 04/10/2018

package models

import (
	"context"
	"errors"
	"io/ioutil"
	"os/exec"
	"time"
)

type Data struct {
	RawData    []byte
	SSConfData []byte
	ResultData []byte
}

//maximum wait time in seconds for stream splitter to process any stream

const CtxTimeout time.Duration = 4

func (data *Data) RunBySS() error {

	//write the input data contents to a file in /var/tmp where SS config will pick the data from.
	if err := ioutil.WriteFile("/var/tmp/testing-input.txt", data.RawData, 0644); err != nil {
		return errors.New("Could not write to an input file")
	}

	//write the plugins config data to plugins.conf. This overrides the existing conf deployed. Application.conf is deployed via ansible and it has include plugins.conf

	if err := ioutil.WriteFile("/etc/stream-splitter/plugins.conf", data.SSConfData, 0644); err != nil {
		return errors.New("Could not write to plugins.conf")
	}

	//input file is ready, now we can start stream splitter.

	//set a context
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout*time.Second)
	defer cancel()

	// set the stream splitter start command.

	cmd := exec.CommandContext(ctx, "service", "stream-splitter", "restart")

	//Read Some SS Log Files

	err := cmd.Run() //Run will immediately return (while the command is running on a fork.)

	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("ERROR: Some issue with starting stream-splitter. SS start or Processing is taking more than 6 sec. I am killing SS and cleaning resources (safety first!).")
	}

	if err != nil {
		//the cmd.Run() returned with non zero exit code.

		return errors.New("ERROR: SS started and exited with non-zero exit code. Is SS running as privileged user?")
	}

	//IF there are no errors then most likely our output file has data that we need to pass back. Let's read that.

	//I hate using sleep here but it appears that Stream Splitter is not flushing the output to file immediately.
	//Reading the file immediately shows no results. 1 sec is fine. But just keeping 2s.
	time.Sleep(2 * time.Second)

	data.ResultData, err = ioutil.ReadFile("/var/tmp/testing-output.txt")

	if err != nil {
		//some problem reading the output
		return errors.New("ERROR: Some issue reading the output file. Check in the server")
	}

	if len(data.ResultData) < 1 {
		return errors.New("WARNING: NO DATA in the OUTPUT file")
	}

	return nil
}
