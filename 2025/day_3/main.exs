defmodule Solve do
  def p2(arr, max) do
    n = length(arr)
    r = 12 - length(max)
    IO.puts("n is: ")
    IO.inspect(n)
    IO.inspect(r)
    IO.inspect(max)
    #    r = n
    cond do
      r == n ->
	Enum.concat(max, arr)
      r == 0 ->
	max
      n > r ->
	maxi = Enum.reduce(Enum.take(arr, n - r + 1), %{cur_idx: -1, max_idx: -1, max: -1},
	  fn val, acc ->
	    if val > acc.max do
	      %{cur_idx: acc.cur_idx + 1, max_idx: acc.cur_idx + 1, max: val}
	    else
	      %{cur_idx: acc.cur_idx + 1, max_idx: acc.max_idx, max: acc.max}
	    end
	  end)
	IO.inspect(maxi)
	if maxi.max_idx < 0 do
	  System.halt(0)
	end
	
	p2(Enum.drop(arr, maxi.max_idx + 1), max ++ [maxi.max])
	#IO.puts("Hello")
	
	#p2([1,2,3,4,5,6],[1,2,3,4,5,6])
    end	
  end
    
  def max_joltage(bank) do
    IO.inspect(bank)
    #Enum.take(bank, 3) |> IO.inspect()
    p2(bank, [])
  end
end

File.read!("input.txt")
|> String.split()
|> Enum.map(& String.to_integer(&1) |> Integer.digits() |> Solve.max_joltage() |> Integer.undigits())
|> Enum.sum()
|> IO.inspect(charlists: :as_lists)
