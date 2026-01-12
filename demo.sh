#!/bin/bash

echo "=== GasLens Bytecode Analyzer Demo ==="
echo

echo "1. Testing with simple bytecode..."
./gaslens bytecode.txt
echo

echo "2. Testing with complex bytecode (with functions and storage)..."
./gaslens test_bytecode.txt
echo

echo "3. Generated files:"
ls -la *.json *.csv 2>/dev/null || echo "No export files found"
echo

echo "4. Sample JSON report (first 10 lines):"
head -10 analysis_report.json 2>/dev/null || echo "No JSON report found"
echo

echo "5. Sample CSV report (first 5 lines):"
head -5 analysis_report.csv 2>/dev/null || echo "No CSV report found"
echo

echo "Demo complete!"
