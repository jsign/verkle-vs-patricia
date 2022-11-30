set title "Storage Tries - Depth"
set key off
set datafile separator ","; 
set xrange [1:11]; 
set xtics 1
set yrange [0:30]; 
set xlabel "depth" offset 1
set ylabel "% of slots at depth" offset 1
set style fill solid; 
set xtics nomirror
set boxwidth 0.5; 
set tic scale 0
set grid ytics
unset border
plot "storagetrie_depth.csv" using 1:2 with boxes