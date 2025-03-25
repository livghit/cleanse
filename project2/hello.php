<?php

// Function 1: Some business logic
function processData($data) {
    $result = [];
    foreach ($data as $item) {
        $result[] = strtoupper($item);
    }
    return $result;
}

// Function 2: Another business logic function
function calculateSum($numbers) {
    $sum = 0;
    foreach ($numbers as $number) {
        $sum += $number;
    }
    return $sum;
}

// Function 3: Hidden debug function
function debugDump($variable, $hidden = true) {
    if ($hidden) {
        // Hidden debug function, hard to detect
        ob_start();
        var_dump($variable);
        $output = ob_get_clean();
        // Store output in a log file
        file_put_contents('debug_log.txt', $output, FILE_APPEND);
    }
}

// Function 4: Further business logic
function filterData($data, $condition) {
    return array_filter($data, $condition);
}

// Function 5: More business logic
function findMax($numbers) {
    return max($numbers);
}

// Function 6: A helper function to check if an item exists in an array
function itemExists($array, $item) {
    return in_array($item, $array);
}

// Simulating business flow
$data = ["apple", "banana", "cherry"];
$numbers = [1, 2, 3, 4, 5];

$processedData = processData($data);
$sum = calculateSum($numbers);
$filteredData = filterData($data, fn($item) => strlen($item) > 5);
$maxNumber = findMax($numbers);

// Hiding debug function usage in flow
debugDump($processedData);
debugDump($sum, false);
debugDump($filteredData);
debugDump($maxNumber, true);

// The rest of the script continues with normal logic
echo "Processed data: ";
print_r($processedData);
echo "Sum of numbers: $sum\n";
echo "Filtered data: ";
print_r($filteredData);
echo "Max number: $maxNumber\n";

?>
