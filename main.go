package main

import(
  "math/rand"
  "fmt"
  "image"
  "image/color"
  "image/gif"
  "os"
  "sandbox/gosorts"
)

func main() {
  c := make(chan gosorts.SortState);
  xs := rand.Perm(50);

  fmt.Println(xs);

  go gosorts.Bubble(xs, c);

  var frames []*image.Paletted;

  for state := range c {
    img := drawArray(state);
    frames = append(frames, img);
  }

  writeGif("result.gif", frames);
}

func drawArray(state gosorts.SortState) *image.Paletted {
  lenVals := len(state.Values);

  w := lenVals;
  h := lenVals;

  black := color.RGBA{0, 0, 0, 255};
  blue := color.RGBA{0, 0, 255, 255};
  red := color.RGBA{255, 0, 0, 255};

  palette := color.Palette{
    black,
    blue,
    red,
  };

  img := image.NewPaletted(image.Rect(0,0,w,h), palette);

  xStep := w / lenVals;
  yStep := h / lenVals;

  colors := make([]color.Color, len(state.Values));

  for i := range state.Values {
    colors[i] = blue;
  }

  for _,i := range state.ConsideringIndices {
    colors[i] = red;
  }

  for i,val := range state.Values {
    for x := i*xStep; x < (i+1)*xStep; x++ {
      for y := h; y > h-(val*yStep); y-- {
        img.Set(x, y, colors[i]);
      }
    }
  }

  return img;
}

func writeGif(filename string, images []*image.Paletted) {
  out, err := os.Create(filename);

  if (err != nil) {
    panic(err);
  }

  defer out.Close();

  delays := make([]int, len(images));

  for i := range delays {
    delays[i] = 10;
  }

  g := gif.GIF{
        Image: images,
        Delay: delays,
  };

  gif.EncodeAll(out, &g);
}

