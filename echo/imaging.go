package main


import (
	"image"
	"image/draw"
	"image/png"
	"github.com/golang/freetype"
	"os"
	"bufio"
	"github.com/bwmarrin/discordgo"
	"bytes"
	"gotools"
	"time"
)








func ImageSystem(s *discordgo.Session, m *discordgo.MessageCreate, cat string, header string, subheader string, font string, font2 string) {
	gto.Print(m.Author.Username + " made a "+cat+" Banner!")
	WriteTeemo(cat, m.Author.ID, 130, 15, font, 40, header)

	testing := subheader
	t := gto.Split(testing, "{br}")
	cnt := 0
	data := ""
	top := 60
	for _, v := range t {
		cnt++
		data = data + v + " "
		WriteTeemo(m.Author.ID, m.Author.ID, 150, top, font2, 13, data)
		top = top + 12
		data = ""
		cnt = 0
	}
	time.Sleep(1000 * time.Millisecond)
	mk, err := gto.ReadFile("images/custom/"+m.Author.ID+".png")
	if err != nil {
		return
	}
	tc := bytes.NewReader(mk)
	_, err = s.ChannelFileSend(m.ChannelID, "EchoTeemoBanner.png", tc)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "I don't have permissions to post images in this channel.")
	}
	os.Remove("images/custom/"+m.Author.ID+".png")
}
















func PostImage(cat string, s *discordgo.Session, m *discordgo.MessageCreate) {
	sky, _ := gto.ReadDir("images/"+cat, "*")
	cnt := len(sky)
	rand := gto.Random(1, cnt)
	mk, err := gto.ReadFile(sky[rand])
	if err == nil {
		s.ChannelTyping(m.ChannelID)
        tc := bytes.NewReader(mk)
        s.ChannelFileSend(m.ChannelID, cat+".png", tc)
	}
}



func WriteFile(input string, output string, left int, top int, thefont string, tsize float64, text string) {

    reader, err := os.Open("images/custom/"+input+".png")
    if err != nil {
        return
    }
    src, _, err := image.Decode(reader)
    if err != nil {
        return
    }


    b := src.Bounds()
    rgba := image.NewRGBA(image.Rect(0, 0, 252, 297))
    draw.Draw(rgba, rgba.Bounds(), src, b.Min, draw.Src)



	// Read the font data.
	fontBytes, err := gto.ReadFile(thefont)
	if err != nil {
	//	log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
	//	log.Println(err)
		return
	}

	// Initialize the context.
	fg := image.Black
	/*
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	*/
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(tsize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)


	// Draw the text.
	pt := freetype.Pt(left, top+int(c.PointToFixed(tsize)>>6))
//	for _, s := range text {
		_, err = c.DrawString(text, pt)
		if err != nil {
		//	log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(tsize * *spacing)
//	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("images/custom/"+output+".png")
	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	}
	defer outFile.Close()
	bk := bufio.NewWriter(outFile)
	err = png.Encode(bk, rgba)
	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	}
	err = bk.Flush()
	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	}
//	fmt.Println("Wrote out.png OK.")
}






func WriteTeemo(input string, output string, left int, top int, thefont string, tsize float64, text string) {

    reader, err := os.Open("images/custom/"+input+".png")
    if err != nil {
    	gto.Print("os.Open() error: ")
    	gto.Print(err)
        return
    }
    src, _, err := image.Decode(reader)
    if err != nil {
    	gto.Print(err)
        return
    }


    b := src.Bounds()
    rgba := image.NewRGBA(image.Rect(0, 0, 600, 100))
    draw.Draw(rgba, rgba.Bounds(), src, b.Min, draw.Src)



	// Read the font data.
	fontBytes, err := gto.ReadFile(thefont)
	if err != nil {
	//	log.Println(err)
		gto.Print(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
	//	log.Println(err)
		gto.Print(err)
		return
	}

	// Initialize the context.
	fg := image.White
	/*
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	*/
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(tsize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)


	// Draw the text.
	pt := freetype.Pt(left, top+int(c.PointToFixed(tsize)>>6))
//	for _, s := range text {
		_, err = c.DrawString(text, pt)
		if err != nil {
		//	log.Println(err)
			gto.Print(err)
			return
		}
		pt.Y += c.PointToFixed(tsize * *spacing)
//	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("images/custom/"+output+".png")
	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
		gto.Print(err)
	}
	defer outFile.Close()
	bk := bufio.NewWriter(outFile)
	err = png.Encode(bk, rgba)
	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
		gto.Print(err)
	}
	err = bk.Flush()
	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	}
//	fmt.Println("Wrote out.png OK.")
	gto.Print(err)
}


