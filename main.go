package main

import (
	"flag"
	"fmt"
	"os"
)

func ReadSecret(path string) (string, error) {
	info, err := os.Stat(path)

	if err != nil {
		if _, ok := err.(*os.PathError); ok { // this isn't necessary
			return path, nil
		} else {
			return "", err
		}
	}

	if info.IsDir() {
		return "", fmt.Errorf("%s is not a file", path)
	}

	key, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Couldn't read key from %s: %s", path, err)
	}

	return string(key), nil
}

func main() {
	dstIP := flag.String("ip", "", "Destination IP")
	secret := flag.String("secret", "", "Secret or File containing Secret")
	port := flag.Uint("port", 0, "Port that should be opened if sequence valid.")

	secretKey, err := ReadSecret(*secret)

	if err != nil {
		fmt.Errorf("Couldn't parse secret: %s", err)
		os.Exit(1)
	}
}
