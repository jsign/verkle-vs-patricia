set title "Storage Tries - Num used slots"
set key off
set datafile separator ","; 
set xrange [0:25]; 
set xtics 1
set yrange [0:65]; 
set xlabel "# non-zero slots" offset 1
set ylabel "% of SCs" offset 1
set style fill solid; 
set boxwidth 0.5; 
set xtics nomirror
set tic scale 0
set grid ytics
unset border
plot "storagetrie_numslots.csv" using 1:2 with boxes