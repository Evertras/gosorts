package gosorts

type SortState struct {
  Values []int;
  ConsideringIndices []int;
}

func Bubble(xs []int, c chan SortState) []int {
  l := len(xs);
  madeSwap := true;
  for i := 0; i < l && madeSwap; i++ {
    madeSwap = false;
    for j,k := 0,1; k < l; j,k = j+1,k+1 {
      considering := []int {j,k};

      c <- SortState{xs,considering};
      modified := make([]int, l);
      copy(modified, xs);
      xs = modified;
      if xs[j] > xs[k] {
        madeSwap = true;
        tmp := xs[j];
        xs[j] = xs[k];
        xs[k] = tmp;
      }
      c <- SortState{xs,considering};
    }
  }

  close(c);
  return xs;
}

