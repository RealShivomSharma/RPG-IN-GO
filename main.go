package main 

import (
    "image/color"
    "image"
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
// Sprite can be more than a player
type Sprite struct {
    Img *ebiten.Image
    X, Y float64
}
// Struct embedding

type Enemy struct {
    *Sprite 
    FollowsPlayer bool
}

type Game struct{
    player *Sprite
    enemies []*Enemy
}

func (g *Game) Update() error {
    // React to keypresses
    if ebiten.IsKeyPressed(ebiten.KeyRight) {
        g.player.X += 2
    }
    if ebiten.IsKeyPressed(ebiten.KeyLeft) {
        g.player.X -= 2
    }

    if ebiten.IsKeyPressed(ebiten.KeyDown) {
        g.player.Y += 2
    }
    if ebiten.IsKeyPressed(ebiten.KeyUp) {
        g.player.Y -= 2
    }

    for _, enemy := range g.enemies{
        if enemy.FollowsPlayer {
            if enemy.X < g.player.X {
                enemy.X += 0.5
            } else if enemy.X > g.player.X {
                enemy.X -= 0.5
            }
            if enemy.Y < g.player.Y {
                enemy.Y += 0.5
            } else if enemy.Y > g.player.Y{
                enemy.Y -= 0.5
            }
    }
    } 
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{120,180,255,255})
    opts := ebiten.DrawImageOptions{}
    opts.GeoM.Translate(g.player.X, g.player.Y)

    // Drawing the player

    screen.DrawImage(g.player.Img.SubImage(
        image.Rect(0,0,32,32),
    ).(*ebiten.Image), &opts)

    opts.GeoM.Reset()
    for _, enemy:= range g.enemies {
        opts.GeoM.Translate(enemy.X, enemy.Y)

        screen.DrawImage(enemy.Img.SubImage(
            image.Rect(0,0,32,32),
        ).(*ebiten.Image), &opts)

        opts.GeoM.Reset()

    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 320,240
}

func main() {
    ebiten.SetWindowSize(640,480)
    ebiten.SetWindowTitle("Hello, World!")
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/Player.png")
    if err != nil {
        log.Fatal(err)
    }
    skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/Skeleton.png")
    if err != nil {
        log.Fatal(err)
    }
    game := Game{
        player: &Sprite{
            Img: playerImg,
            X: 100,
            Y: 100,
        },
        enemies : []*Enemy{
            {
                Sprite: &Sprite{
                    Img: skeletonImg,
                    X: 50,
                    Y: 50,
                },
                FollowsPlayer: true,
            },
            {
                Sprite: &Sprite{
                    Img: skeletonImg,
                    X: 150,
                    Y: 50,
                },
                FollowsPlayer: false,
            },
            {
                Sprite: &Sprite{
                    Img: skeletonImg,
                    X: 100,
                    Y: 150,
                },
                FollowsPlayer: true,
            },
        },
    }
    if err := ebiten.RunGame(&game); err != nil {
        log.Fatal(err)
    }
}



