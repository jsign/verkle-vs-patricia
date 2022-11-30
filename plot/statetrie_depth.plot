set title "State Trie - Depth"
set key off
set datafile separator ","; 
set xrange [8.5:12]; 
set xtics nomirror 1
set yrange [0:55]; 
set xlabel "depth" offset 1
set ylabel "% of addresses at depth" offset 1
set style fill solid; 
set boxwidth 0.5; 
set tic scale 0
set grid ytics
unset border
plot "statetrie_depth.csv" using 1:2 with boxes