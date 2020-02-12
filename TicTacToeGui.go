package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/gfx"

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("TicTacToe GUI", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		593, 593, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	renderer,err := sdl.CreateSoftwareRenderer(surface)
	
	if err != nil {
		panic(err)
	}

	running := true
	lastx,lasty:=int32(0),int32(0)
	mouseDown:=false
	xyCounter:=0;
	var alphax uint8 =0;
	var alphax1 uint8 =0;
	var alphay uint8 =0;
	
	for running {
		//Draw #
		window.UpdateSurface()
		//Draw Board
		surface.FillRect(nil, 0xffffff)
		//Section 1
		//Square
		square1:=&sdl.Rect{0, 0, 193, 193}
		surface.FillRect(square1, 0x000000)
		if xyCounter==0 {
			if mouseDown {
				alphax1=255
			}
			//xyCounter=1
		}
		gfx.ThickLineRGBA(renderer, 0, 0, 193, 193, 20,0,255,0,alphax1)
		gfx.ThickLineRGBA(renderer, 0, 193, 193, 0, 20,0,255,0,alphax1)
		//X
		
		//Section 2
		//Square
		square2:=&sdl.Rect{200, 0, 193, 193}
		surface.FillRect(square2, 0x000000)
		
		//Section 3
		//Square
		square3:=&sdl.Rect{400, 0, 193, 193}
		surface.FillRect(square3, 0x000000)
		
		square4:=&sdl.Rect{0, 200, 193, 193}
		surface.FillRect(square4, 0x000000)
		square5:=&sdl.Rect{200, 200, 193, 193}
		surface.FillRect(square5, 0x000000)
		square6:=&sdl.Rect{400, 200, 193, 193}
		surface.FillRect(square6, 0x000000)
		
		square7:=&sdl.Rect{0, 400, 193, 193}
		surface.FillRect(square7, 0x000000)
		square8:=&sdl.Rect{200, 400, 193, 193}
		surface.FillRect(square8, 0x000000)
		square9:=&sdl.Rect{400, 400, 193, 193}
		surface.FillRect(square9, 0x000000)
	
			
		
		//loop through struct and draw placed pieces
		//draw phantom piece if mouse is over blank
		
		//Hover Box
		// sdl.Rect(x pos, y pos, width, height)
		//rect:=&sdl.Rect{lastx/200*200, lasty/200*200, 193, 193}
		//rect:=&sdl.Rect{lastx/200*200, lasty/200*200, 193, 193}
		if xyCounter==0 {
			alphax=100
			alphay=0
			if mouseDown {
				alphax=255
			}
			//xyCounter=1
		} else if xyCounter==1{
			alphax=0
			alphay=100
			if mouseDown {
				alphay=255
			}
			xyCounter=0
		}
		gfx.ThickLineRGBA(renderer, lastx+0, lasty+0, lastx+193, lasty+193, 20,0,255,0,alphax)
		gfx.ThickLineRGBA(renderer, lastx+0, lasty+193, lastx+193, lasty+0, 20,0,255,0,alphax)
		gfx.CircleRGBA(renderer, lastx+0, lasty+0, 96, 0,255,0,alphay)
		//Hover Box Color
		//color:=uint32(0x7fff00)
		
		
		//if win cond then do a thing
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			//~ window.UpdateSurface()
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.MouseMotionEvent:
				lastx,lasty=e.X,e.Y
			case *sdl.MouseButtonEvent:
				if e.Button==sdl.BUTTON_LEFT && e.State==sdl.PRESSED {
					mouseDown=true
				}
				if e.Button==sdl.BUTTON_LEFT && e.State==sdl.RELEASED {
					mouseDown=false
				}
				//trigger place
				//if in win state reset
			} 
		}
	}
}
