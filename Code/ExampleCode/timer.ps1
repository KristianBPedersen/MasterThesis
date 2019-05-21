$n = 100 #Number of experiment repetitions
$outliersSize = 2 #Number of outliers to remove

$times  = @()

#Running the experiment n times
for ($i = 0; $i -lt $n; $i++){
    $time = Measure-Command -Expression {./composite.exe}
    $times += $time.TotalMilliseconds
}

#Sorting to remove outliers
$times = $times | Sort

#Summing all the times except the outliers
$total = 0
for ($i = $outliersSize; $i -lt $n - $outliersSize; $i++){
    $total += $times[$i]
}

#Printing the average time
echo ($total/($n - $outliersSize))