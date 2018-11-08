#!/usr/local/bin/julia

if length(ARGS) < 2
  println("need 2 arguments.")
  exit(1)
end

x = parse(Int, ARGS[1])
y = parse(Int, ARGS[2])
total = x + y
println(total)
