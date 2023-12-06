# Monitor X

Monitor-X is a system monitoring tool that provides real-time insights into various aspects of system performance, including CPU usage, memory utilization, disk usage, network statistics, and more. Built with Go and termui, it offers a visually appealing and easy-to-navigate terminal-based user interface.

## Features

- **CPU Monitoring**: Track CPU usage per core, total CPU usage, and top CPU-consuming processes.
- **Memory Usage**: Monitor total, used, and free memory in real-time.
- **Disk Usage**: Keep an eye on disk space utilization.
- **Network Statistics**: View detailed network interface statistics, including bytes sent/received.
- **Load Average**: Check system load average over 1, 5, and 15 minutes.
- **User-friendly Interface**: Navigable terminal-based UI for easy monitoring.



## Installation

To install Monitor X, ensure that you have Golang installed on your system. Follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/EmaLinuxawy/monitor-x.git
   ```
2. Navigate to the cloned directory:

   ```bash
   cd monitor-x
   ```
3. Build the application:

   ```bash
   go build .
   ```

4. Run the application:

   ```bash
   ./monitor-x
   ```
   
## Usage

After starting Monitor-X, use the following keyboard shortcuts to navigate through the application:

- `q` or `<Ctrl-C>`: Quit the application.
- `<Up>` and `<Down>`: Scroll through the lists.
- `<Enter>`: Expand or collapse sections (if applicable).

## Contributing

Contributions to Monitor-X are welcome! Please feel free to submit pull requests, report bugs, or suggest features through the [GitHub repository](https://github.com/Emalinuxawy/monitor-x).
