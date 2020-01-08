# 🍾 jingo

`jingo` is a [Jinja2](jinja-docs) string interpolation library for Go. This allows Go-project owners some flexability in their templating and string interpolation, especially if they themselves or other contributors come from a Python background. It attempts to duplication as much of the functionality and configuration of the original library.

**This project is in a VERY infant stage right now. I, Dan, would like to set up the internals before accepting major contributions, so please check back if you are interested in using jingo.**

## Getting Started

If you want to use `jinjo` in your own project, you can add it with the following command.

```bash
go get github.com/dstarner/jingo
```

## Development

If you want to work on `jinjo` yourself, feel free to clone this repository to your local machine. Go 1.12 or 1.13 is required. Once set up, most of the common commands can be executed through the provided `Makefile`. Check out this file to get the basic targets that can be run.

[jinja-docs]: https://jinja.palletsprojects.com/en/2.10.x/
