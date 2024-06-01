## GoLang App README.md

### Project Overview

This GoLang application provides a command-line tool for managing configuration settings in a YAML file named `.config.yaml`. It offers four main functionalities:

1. **Add:** Add new key-value pairs to the configuration file.

2. **Update:** Modify existing key-value pairs in the configuration file.

3. **Delete:** Remove specific key-value pairs from the configuration file.

4. **List:** Display all currently defined key-value pairs in the configuration file.

### Installation

1. **Prerequisites:** Ensure you have GoLang installed on your system. You can download and install GoLang from the official website: [https://go.dev/doc/install](https://go.dev/doc/install)

2. **Cloning the Repository:** Clone this repository to your local machine using Git:

   ```bash
   git clone <repository_url>
   ```

3. **Building the Application:** Navigate to the project directory and build the executable file:

   ```bash
   go build
   ```

### Usage

1. **Running the Application:** Execute the built executable file followed by the desired operation and its corresponding arguments:

   ```bash
   ./mycobra [operation] [flags]
   ```

2. **Operation Options:**

   - `add`: Add a new key-value pair.

   - `update`: Update an existing key-value pair.

   - `delete`: Remove a specific key-value pair.

   - `list`: Display all key-value pairs.

3. **Key and Value Arguments:**

   - `key`: The name of the configuration setting to be modified (required for all operations).

   - `value`: The value associated with the configuration setting (required for `add` and `update` operations).

### Examples

1. **Adding a Key-Value Pair:**

   ```bash
   ./mycobra add -k server_port -v 8080
   ```

2. **Updating a Key-Value Pair:**

   ```bash
   ./mycobra update -k server_port -v 9000
   ```

3. **Deleting a Key-Value Pair:**

   ```bash
   ./mycobra delete -k database_name
   ```

4. **Listing All Key-Value Pairs:**

   ```bash
   ./mycobra list
   ```
