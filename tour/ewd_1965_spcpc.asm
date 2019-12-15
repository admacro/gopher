;; Solution of a problem in concurrent programming control
;; Edsger W. Dijkstra
;; 1965

;; Boolean array b, c[1:N]; int k (1 <= k <= N)
;; b and c are set to true; the starting value of k is immaterial

int j
Li0: b[i] = false
Li1: if k != i then
Li2:   begin c[i] = true
Li3:     if b[k] then k = i
         go to Li1
       end
     else
Li4:   begin c[i] = false
         for j = 1 step 1 until N do
           if j != i and !c[j] then go to Li1
       end
       ;; critical section
       c[i] = true
       b[i] = true
       ;; remainder of the cycle in which stopping is allowed
       go to Li0


;; N = 3
;; i 1
;; b[1] false true
;; c[1] false true
;; c[2] true
;; c[3] true
;; k 1
;; j 1 2 3
;; critical section

;; i 2
;; b[2] false true
;; c[1] true
;; c[2] true true false true
;; c[3] true
;; k 1 2
;; j 1 2 3
;; critical section
