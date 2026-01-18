# anonlookerGoLang

A collection of simple desktop GUI applications built with Go and the Fyne framework.

## Projects

### 1. Stopwatch

A stopwatch application with millisecond precision.

**Features:**
- Start/Stop/Reset controls
- Displays time in HH:MM:SS.MS format
- Concurrent timing using goroutines
- User notifications for invalid actions

**Location:** `stopwitch/`

### 2. Guess Prime Number

An interactive game where you guess whether a randomly generated number is prime.

**Features:**
- Random number generation with configurable range
- Settings window to adjust min/max values
- Instant feedback on your guesses
- Simple primality testing algorithm

**Location:** `GuessPrimeNumber/`

## Requirements

- Go 1.16 or later
- Fyne v2.7.1

## Installation

Each application is self-contained with its own `go.mod` file.

### Install Stopwatch

```bash
cd stopwitch
go build
./stopwitch
```

### Install Guess Prime Number

```bash
cd GuessPrimeNumber
go build
./GuessPrimeNumber
```

## Usage

### Stopwatch
1. Click **Start** to begin timing
2. Click **Stop** to pause
3. Click **Reset** to clear the timer

### Guess Prime Number
1. A random number is displayed
2. Click **Yes** if you think it's prime, **No** if not
3. You'll get instant feedback and a new number
4. Click the settings icon (⚙️) to adjust the number range

## Technology Stack

- **Language:** Go
- **GUI Framework:** [Fyne](https://fyne.io/) v2.7.1
- **Concurrency:** Goroutines and Mutexes

## License

See the [LICENSE](LICENSE) file for details.
