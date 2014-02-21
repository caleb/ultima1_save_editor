package main

import(
  "fmt"
  "bufio"
  "os"
  "io"
  "encoding/binary"
  "flag"
  "strconv"
  "strings"
  "github.com/wsxiaoys/terminal/color"
)

func bytesToUint16LE(b1 byte, b2 byte) (i uint16) {
  i = binary.LittleEndian.Uint16([]byte{b1, b2})
  return
}

func bytesToUint32LE(b1 byte, b2 byte, b3 byte, b4 byte) (i uint32) {
  i = binary.LittleEndian.Uint32([]byte{b1, b2, b3, b4})
  return
}

func readLine() (s string, err error) {
  bio := bufio.NewReader(os.Stdin)

  var hasMoreLine bool
  var bytes []byte

  bytes, hasMoreLine, err = bio.ReadLine()
  if err != nil { return }

  for hasMoreLine {
    var bytes2 []byte

    bytes2, hasMoreLine, err = bio.ReadLine()
    if err != nil { return }

    bytes = append(bytes, bytes2...)
  }

  s = string(bytes)

  return
}

func readWholeNumber(prompt string, bits int, unsigned bool, allowBlank bool, defaultValue int64) (out int64, err error) {
  var iu uint64
  var i int64

  loop: for {
    fmt.Print(prompt)

    line, err := readLine()

    if strings.TrimSpace(line) == "" && allowBlank {
      if unsigned {
        iu = uint64(defaultValue)
      } else {
        i = defaultValue
      }

      break loop
    }

    if unsigned {
      iu, err = strconv.ParseUint(string(line), 10, bits)
    } else {
      i, err = strconv.ParseInt(string(line), 10, bits)
    }

    if err != nil {
      numError := err.(*strconv.NumError)

      switch numError.Err {
      case strconv.ErrSyntax:
        if unsigned {
          color.Printf("@{r}%s is not a non-negative integer\n", numError.Num)
        } else {
          color.Printf("@{r}%s is not a integer\n", numError.Num)
        }
        continue loop
      case strconv.ErrRange:
        color.Printf("@{r}%s does not fit in %d bits\n", numError.Num, bits)
        continue loop
      default:
        color.Printf("@{r}Unknown error %s\n", numError)
        continue loop
      }
    }

    break loop;
  }

  if unsigned {
    out = int64(iu)
  } else {
    out = i
  }

  return
}

func readUint16(prompt string, allowBlank bool, defaultValue uint16) (out uint16, err error) {
  i, err := readWholeNumber(prompt, 16, true, allowBlank, int64(defaultValue))

  if err == nil {
    out = uint16(i)
  }

  return
}

func readUint32(prompt string, allowBlank bool, defaultValue uint32) (out uint32, err error) {
  i, err := readWholeNumber(prompt, 32, true, allowBlank, int64(defaultValue))

  if err == nil {
    out = uint32(i)
  }

  return
}

func main() {
  outputFilename := flag.String("o", "", "The file to write the modified save file to")
  flag.Parse()

  inputFilename := flag.Arg(0)

  if *outputFilename == "" || inputFilename == "" {
    fmt.Println("usage: u1-save-editor -o <output file> <ultima 1 save file>")
    os.Exit(-1)
  }

  saveFile, err := os.Open(inputFilename)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  defer saveFile.Close()

  fi, err := saveFile.Stat();

  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  if ! fi.Mode().IsRegular() {
    fmt.Printf("\"%v\" is not a regular file")
    os.Exit(-1)
  }

  fmt.Printf("Opening save file %v\n\n", flag.Arg(0))

  bufferedReader := bufio.NewReader(saveFile)

  var name [14]byte
  var strength uint16
  var agility uint16
  var stamina uint16
  var charisma uint16
  var wisdom uint16
  var intelligence uint16
  var hitPoints uint16
  var gold uint16
  var experience uint16
  var food uint16
  var moves uint32

  saveFileBytes := make([]byte, 0)

  loop: for {
    b, err := bufferedReader.ReadByte()
    switch err {
    case nil:
      saveFileBytes = append(saveFileBytes, b)
    case io.EOF:
      break loop
    }
  }

  copy(name[:], saveFileBytes[:14])
  strength = bytesToUint16LE(saveFileBytes[24], saveFileBytes[25])
  agility = bytesToUint16LE(saveFileBytes[26], saveFileBytes[27])
  stamina = bytesToUint16LE(saveFileBytes[28], saveFileBytes[29])
  charisma = bytesToUint16LE(saveFileBytes[30], saveFileBytes[31])
  wisdom = bytesToUint16LE(saveFileBytes[32], saveFileBytes[33])
  intelligence = bytesToUint16LE(saveFileBytes[34], saveFileBytes[35])

  hitPoints = bytesToUint16LE(saveFileBytes[22], saveFileBytes[23])
  gold = bytesToUint16LE(saveFileBytes[36], saveFileBytes[37])
  food = bytesToUint16LE(saveFileBytes[40], saveFileBytes[41])
  experience = bytesToUint16LE(saveFileBytes[38], saveFileBytes[39])

  moves = bytesToUint32LE(saveFileBytes[172], saveFileBytes[173], saveFileBytes[174], saveFileBytes[175])


  fmt.Printf("name (%v): ", string(name[:]))
  line, err := readLine()
  if err != nil { return }

  if strings.TrimSpace(line) != "" {
    for i, _ := range line {
      if i == 14 { break }
      name[i] = line[i]
    }
  }

  strength, _     = readUint16(fmt.Sprintf("strength (%v): ", strength), true, strength)
  agility, _      = readUint16(fmt.Sprintf("agility (%v): ", agility), true, agility)
  stamina, _      = readUint16(fmt.Sprintf("stamina (%v): ", stamina), true, stamina)
  charisma, _     = readUint16(fmt.Sprintf("charisma (%v): ", charisma), true, charisma)
  wisdom, _       = readUint16(fmt.Sprintf("wisdom (%v): ", wisdom), true, wisdom)
  intelligence, _ = readUint16(fmt.Sprintf("intelligence (%v): ", intelligence), true, intelligence)
  hitPoints, _    = readUint16(fmt.Sprintf("hit points (%v): ", hitPoints), true, hitPoints)
  food, _         = readUint16(fmt.Sprintf("food (%v): ", food), true, food)
  experience, _   = readUint16(fmt.Sprintf("experience (%v): ", experience), true, experience)
  gold, _         = readUint16(fmt.Sprintf("gold (%v): ", gold), true, gold)
  moves, _        = readUint32(fmt.Sprintf("moves (%v): ", moves), true, moves)

  var bytes [4]byte

  // name
  copy(saveFileBytes[0:14], name[:])

  // strength
  binary.LittleEndian.PutUint16(bytes[:3], strength)
  copy(saveFileBytes[24:26], bytes[:])

  // agility
  binary.LittleEndian.PutUint16(bytes[:3], agility)
  copy(saveFileBytes[26:28], bytes[:])

  // stamina
  binary.LittleEndian.PutUint16(bytes[:3], stamina)
  copy(saveFileBytes[28:30], bytes[:])

  // charisma
  binary.LittleEndian.PutUint16(bytes[:3], charisma)
  copy(saveFileBytes[30:32], bytes[:])

  // wisdom
  binary.LittleEndian.PutUint16(bytes[:3], wisdom)
  copy(saveFileBytes[32:34], bytes[:])

  // intelligence
  binary.LittleEndian.PutUint16(bytes[:3], intelligence)
  copy(saveFileBytes[34:36], bytes[:])

  // hitPoints
  binary.LittleEndian.PutUint16(bytes[:3], hitPoints)
  copy(saveFileBytes[22:24], bytes[:])

  // gold
  binary.LittleEndian.PutUint16(bytes[:3], gold)
  copy(saveFileBytes[36:38], bytes[:])

  // experience
  binary.LittleEndian.PutUint16(bytes[:3], food)
  copy(saveFileBytes[38:40], bytes[:])

  // food
  binary.LittleEndian.PutUint16(bytes[:3], food)
  copy(saveFileBytes[40:42], bytes[:])

  // moves
  binary.LittleEndian.PutUint32(bytes[:], moves)
  copy(saveFileBytes[172:176], bytes[:])

  // Write the output
  outFile, err := os.OpenFile(*outputFilename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)

  if err != nil {
    fmt.Printf("\"%v\" cannot be opened for writing\n", *outputFilename)
    os.Exit(-1)
  }

  fmt.Printf("\n\nOpening \"%v\"\n\n", *outputFilename)

  out := bufio.NewWriter(outFile)
  out.Write(saveFileBytes)
  out.Flush()
}

