package internal

import (
	"fmt"

	testerutils "github.com/codecrafters-io/tester-utils"
)

func antiCheatBasic(stageHarness *testerutils.StageHarness) error {
	b := NewHTTPServerBinary(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	logger := stageHarness.Logger

	client := NewHTTPClient()

	resp, err := client.Get(URL)
	if err != nil {
		return nil
	}

	if resp.Proto != "HTTP/1.1" {
		return fail(logger)
	}

	if date := resp.Header.Get("Date"); date != "" {
		return fail(logger)
	}

	if server := resp.Header.Get("Server"); server != "" {
		return fail(logger)
	}

	return nil
}

func fail(logger *testerutils.Logger) error {
	logger.Criticalf("anti-cheat (ac1) failed.")
	logger.Criticalf("Are you sure you aren't running this against an actual HTTP server?")
	return fmt.Errorf("anti-cheat (ac1) failed")
}
