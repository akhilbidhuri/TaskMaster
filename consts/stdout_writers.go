package consts

import (
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

var Output = colorable.NewColorableStdout()

var RedBold = color.New(color.FgRed).Add(color.Bold)

var Red = color.New(color.FgRed)

var Green = color.New(color.FgGreen)

var Blue = color.New(color.FgBlue)

var BlueBold = color.New(color.FgBlue).Add(color.Bold)
