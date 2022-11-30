set title "Storage Tries - Num non-zero slots"
set key off
set datafile separator ","; 
set xrange [0:40]; 
set xtics 2
set yrange [0:45]; 
set xlabel "# non-zero slots" offset 1
set ylabel "% of SCs" offset 1
set style fill solid; 
set boxwidth 0.5; 
plot "storagetrie_numslots.csv" using 1:2 with boxes