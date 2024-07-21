#!/bin/bash

# Function to clean up background processes
cleanup() {
    echo "Cleaning up..."
    kill $SERVER_PID
    kill $PROFILE_PID
    exit 0
}

# Trap SIGINT (CTRL-C) to call cleanup
trap cleanup SIGINT

# Start the server in the background
make run &
SERVER_PID=$!
sleep 5

# Start the profiler in the background
make profile &
PROFILE_PID=$!

# Wait a bit to ensure server and profiler are up
sleep 5

# Run benchmarks 4 times
for i in {1..4}
do
    make benchmark
done

# Wait for background processes (this keeps the script running)
wait $SERVER_PID
wait $PROFILE_PID
