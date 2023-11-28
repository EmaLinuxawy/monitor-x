# Monitor X

## Introduction

Monitor X is a practice for my knowledge to monitor various system metrics developed in Golang.

## Features

Monitor X provides real-time insights into various aspects of system performance, including:

- Load Average
- Total CPU Usage
- Individual CPU Percentages
- Memory Usage

This tool is built to be lightweight, efficient, and easy to integrate into existing monitoring solutions.

## Installation

To install Monitor X, ensure that you have Golang installed on your system. Follow these steps:

1. Clone the repository:

   ```shell
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

## Usage

```bash
./monitor-x
```

### Functions

1. **getLoadAverage** :
   Retrieves the system's load average, a measure of system activity, giving an idea of how busy the system is.
2. **getTotalCPUUsage** :
   Calculates the total CPU usage percentage, providing a quick overview of CPU utilization.
3. **getCPUPercentages** :
   Lists the usage percentage of each CPU, helpful for identifying uneven load distribution across CPUs.
4. **printCPUstats** :
   Prints detailed CPU statistics in a human-readable format, ideal for logging and monitoring dashboards.
5. **getMemoryUsed** :
   Reports the amount of memory currently being used, essential for detecting memory leaks or high memory consumption.
6. **main** :
   The entry point of the application, orchestrating the monitoring processes and outputting the data.
