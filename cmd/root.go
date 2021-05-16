/*
Copyright © 2021 Zoraiz Hassan <hzoraiz8@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strings"

	imgMani "github.com/TheZoraiz/ascii-image-converter/image_manipulation"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	compl      bool
	dimensions []int
	save       bool

	rootCmd = &cobra.Command{
		Use:   "ascii-image-converter [image path]",
		Short: "Converts images into ascii format",
		Long:  `ascii-image-converter converts images into ascii format and prints them onto the terminal window. Further configuration can be managed with flags`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			numberOfDimensions := len(dimensions)
			if dimensions != nil && numberOfDimensions != 2 {
				return fmt.Errorf("-d requires two dimensions got %v", numberOfDimensions)
			}

			imagePath := args[0]
			return convertPicture(imagePath, compl, dimensions, save)
		},
	}
)

func convertPicture(imagePath string, isComplex bool, dimensions []int, save bool) error {
	pic, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("Unable to open file: %w", err)
	}
	defer pic.Close()

	imData, _, err := image.Decode(pic)
	if err != nil {
		return fmt.Errorf("Unable to decode file: %w", err)
	}

	imgSet := imgMani.ConvertToTerminalSizedSlices(imData, dimensions)
	var asciiSet [][]string

	if isComplex {
		asciiSet = imgMani.ConvertToAsciiDetailed(imgSet)
	} else {
		asciiSet = imgMani.ConvertToAsciiSimple(imgSet)
	}

	ascii := flattenAscii(asciiSet)
	for _, line := range ascii {
		fmt.Println(line)
	}

	if save {
		return ioutil.WriteFile("ascii-image.txt", []byte(strings.Join(ascii, "\n")), 0777)
	}
	return nil
}

// flattenAscii flattens a two-dimensional grid of ascii characters into a one dimension
// of lines of ascii
func flattenAscii(asciiSet [][]string) []string {
	var ascii []string
	for _, line := range asciiSet {
		ascii = append(ascii, strings.Join(line, ""))
	}
	return ascii
}

// Cobra configuration from here on

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-image-converter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&compl, "complex", "c", false, "Prints ascii characters in a larger range, may result in higher quality")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length e.g. 100,30 (defaults to terminal size)")
	rootCmd.PersistentFlags().BoolVarP(&save, "save", "S", false, "Save ascii text in current directory in an ascii-image.txt file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ascii-image-converter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ascii-image-converter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
