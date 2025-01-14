# Flexion Coding Challenge

## ABOUT

This is a unit conversion tool designed for teachers to grade student worksheets. The user provides a worksheet in standard format and student responses in standard format and the program outputs the results to a specified file location.

### Usage

#### Allowed Units
- <b>Temperatures</b>: Kelvin, Celsius, Fahrenheit, Rankine
- <b>Volumes</b>: Liters, Gallons, Cups, Tablespoons, Cubic Inches, Cubic Feet

#### Worksheet Example
Do not include headers on input
| Input | From Unit     | To Unit        |
|-------|---------------|----------------|
| 100   | Fahrenheit    | Rankine        |
| 100   | Kelvin        | Celsius        |
| 100   | Liters        | Gallons        |

#### Response Example (for worksheet above)
Do not include headers on input
|Student Name|Question 1|Question 2|Question 3|
|---------------|-----------|-----------|------|
| Doug Doenges | 310.9   | -173.2  | 26.4 |
| Student B    | 310.928 | -173.15 | 28   |

#### Example Results file output (from worksheet and responses above)
| Input | From Unit     | To Unit        | Correct Answer| | Doug Doenges |           | Student B |           |
|-------|---------------|----------------|---------------|-|--------------|-----------|-----------|-----------|
| 100   | Fahrenheit    | Rankine        | 310.9         | | 310.9        | Correct   | 310.928   | Correct
| 100   | Kelvin        | Celsius        | -173.2        | | -173.2       | Correct   | -173.15   | Correct
| 100   | Liters        | Gallons        | 26.4          | | 26.4         | Correct   | 28        | Incorrect |

## Installation Steps

1. Download the latest release from the GitHub [Releases](https://github.com/dougdoenges/flexion-coding-challenge/releases) page. Choose the relevant executable for your system.
2. For Linux, run the command ```chmod +x flexion-coding-challenge-linux```
3. For MacOS, run the command ```chmod +x flexion-coding-challenge-macos```
4. If your device interprets the executable as malware, open privacy settings and allow this application to be opened.

## How to run the application

```sh
./flexion-coding-challenge-{distribution} --worksheet={path to worksheet file} --responses={path to response file} --output={path and to desired file output}
```

## Prioritized list of development tasks
1. Add help options for end users to receive example file formats for usage and more
2. Deploy packaged code with CI/CD so the project can be used globally on download
3. Increase test coverage for various file inputs
4. Add support for configurable local execution options
5. Add option to provide combinations of input/output via terminal in place of files
6. Add standardized logging package and log during execution
7. Automate semantic versioning for new releases
8. Depending on future requirements, adopt comprehensive unit conversion package [go-units](https://pkg.go.dev/github.com/bcicen/go-units)