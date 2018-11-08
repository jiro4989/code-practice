#!/usr/local/bin/julia

using Printf

if length(ARGS) < 1
  println("need 2 arguments.")
  exit(1)
end

open(ARGS[1], "r") do fp
  i = 0
  for line in eachline(fp)
    i = i + 1
    @printf "%03d:%s\n" i line
  end
end
