set title "Storage Tries - Depth"
set key off
set datafile separator ","; 
set xrange [0:11]; 
set xtics 1
set yrange [0:25]; 
set xlabel "depth" offset 1
set ylabel "% of slots at depth" offset 1
set style fill solid; 
set boxwidth 0.5; 
plot "storagetrie_depth.csv" using 1:2 with boxes