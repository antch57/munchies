# munchies :pizza:

**Munchies** is a CLI tool designed to help you track the snacks you munch on during your workday. Whether you're a casual grazer or a dedicated snacker, Munchies has got you covered!

## Features
- :bar_chart: **Snack Tracking**: Log every snack you eat with ease.

## Why Munchies?
Because snacking is serious business! Munchies helps you stay mindful of your habits while adding a touch of fun to your workday.

---

## Prerequisites

Before installing Munchies, ensure you have the following installed on your system:

- **Go (Golang)**: Munchies requires Go to build and run. You can download it from [golang.org](https://golang.org).
- **Make**: Used for building the project. Most Unix-based systems come with `make` pre-installed.

---

## Installation

Ready to take control of your snack game? Install Munchies and start tracking today! :rocket: :pizza:

1. Clone the repository:
   ```shell
   git clone https://github.com/antch57/munchies.git
   ```

1. Navigate to the project directory:
   ```shell
   cd munchies
   ```

1. Build the CLI tool:
   ```shell
   make install
   ```
    > :information_source: **Note:** This tool will be installed to `/usr/local/bin/`

1. Run the tool:
   ```shell
   munchies --help
   ```

---

## Basic Usage

Start tracking your snacks with simple commands:

1. **Log a snack**:
   ```shell
   munchies add --snack chip --count 1
   ```

    > :information_source: **Note**: Data saved will be in $HOME/.munchies/data/snacks.json

1. **View your snack history**:
   ```shell
   munchies list
   ```

---

Enjoy snacking responsibly with **munchies**! :pizza:
