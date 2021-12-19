package main

import "fmt"

type Decoder struct {
	Signals    []*Signal
	DecoderMap map[string]int
}

func (d *Decoder) AddSignal(s *Signal) {
	d.Signals = append(d.Signals, s)

	for _, signal := range s.Output {
		switch len(signal) {
		case 2:
			val, ok := d.DecoderMap[signal]
			if ok {
				if val != 1 {
					panic("double input")
				}
			} else {
				d.DecoderMap[signal] = 1
			}
		case 3:
			val, ok := d.DecoderMap[signal]
			if ok {
				if val != 7 {
					panic("double input")
				}
			} else {
				d.DecoderMap[signal] = 7
			}
		case 4:
			val, ok := d.DecoderMap[signal]
			if ok {
				if val != 4 {
					panic("double input")
				}
			} else {
				d.DecoderMap[signal] = 4
			}
		case 7:
			val, ok := d.DecoderMap[signal]
			if ok {
				if val != 8 {
					panic("double input")
				}
			} else {
				d.DecoderMap[signal] = 8
			}
		}
	}
}

func (d *Decoder) PrintScrambled() {
	for _, sig := range d.Signals {
		str := ""
		for _, text := range sig.Input {
			if str == "" {
				str = text
			} else {
				str = fmt.Sprintf("%s %s", str, text)
			}
		}

		str = fmt.Sprintf("%s |", str)

		for _, text := range sig.Output {
			str = fmt.Sprintf("%s %s", str, text)
		}
		fmt.Println(str)
	}
}

func (d *Decoder) PrintDecoded() {
	for _, sig := range d.Signals {
		str := ""
		for _, text := range sig.Input {
			if str == "" {
				val, ok := d.DecoderMap[text]
				if ok {
					str = fmt.Sprintf("%d", val)
				} else {
					str = text
				}
			} else {
				val, ok := d.DecoderMap[text]
				if ok {
					str = fmt.Sprintf("%s %d", str, val)
				} else {
					str = fmt.Sprintf("%s %s", str, text)
				}
			}
		}

		str = fmt.Sprintf("%s |", str)

		for _, text := range sig.Output {
			if str == "" {
				val, ok := d.DecoderMap[text]
				if ok {
					str = fmt.Sprintf("%d", val)
				} else {
					str = text
				}
			} else {
				val, ok := d.DecoderMap[text]
				if ok {
					str = fmt.Sprintf("%s %d", str, val)
				} else {
					str = fmt.Sprintf("%s %s", str, text)
				}
			}
		}

		fmt.Println(str)
	}
}

func (d *Decoder) Part1() {
	sum := 0
	for _, sig := range d.Signals {
		for _, text := range sig.Output {
			_, ok := d.DecoderMap[text]
			if ok {
				sum += 1
			}
		}

	}
	fmt.Printf("total: %d\n", sum)
}
