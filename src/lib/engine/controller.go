package engine

import (
	"fmt"
	"rpg-game/src/lib/components"
	"rpg-game/src/lib/config"
	"rpg-game/src/lib/entities"
	"rpg-game/src/lib/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Controller struct{}

func (c *Controller) MovePlayer(player *entities.Player, worldColliders *components.Colliders) {
	player.ResetVelocity()
	player.UpdateCooldowns()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.SetChangeInY(-2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		player.SetChangeInY(2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.SetChangeInX(-2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.SetChangeInX(2)
	}

	player.AddX(player.GetChangeInX())
	CheckCollisionHorizontal(player, worldColliders)

	player.AddY(player.GetChangeInY())
	CheckCollisionVertical(player, worldColliders)

	player.SetCurrentAnimation(
		int(player.GetChangeInX()),
		int(player.GetChangeInY()),
	)
	player.UpdateCurrentAnimation()
}

func (c *Controller) MoveEnemies(
	player *entities.Player,
	enemies map[uint16]*entities.Enemy,
	worldColliders *components.Colliders,
	controls *Controls,
	camera *render.Camera,
) {
	cursorX, cursorY := controls.Mouse.GetMousePositionRelativeToCamera(camera)
	clicked := controls.Mouse.GetMouseButtonState(ebiten.MouseButtonLeft)

	for enemyIndex, enemy := range enemies {
		enemy.UpdateCooldowns()
		enemy.ResetVelocity()

		if enemy.ShouldFollowPlayer {
			if enemy.GetX() < player.GetX() {
				enemy.SetChangeInX(1)
			}
			if enemy.GetX() > player.GetX() {
				enemy.SetChangeInX(-1)
			}
			if enemy.GetY() < player.GetY() {
				enemy.SetChangeInY(1)
			}
			if enemy.GetY() > player.GetY() {
				enemy.SetChangeInY(-1)
			}
		}

		enemy.AddX(enemy.GetChangeInX())
		CheckCollisionHorizontal(enemy, worldColliders)

		enemy.AddY(enemy.GetChangeInY())
		CheckCollisionVertical(enemy, worldColliders)

		if cursorX > enemy.Collider.Min.X && cursorX < enemy.Collider.Max.X && cursorY > enemy.Collider.Min.Y && cursorY < enemy.Collider.Max.Y {
			distanceFromPlayer := player.GetDistanceFrom(enemy)
			if clicked && distanceFromPlayer < config.DefaultTileSizeInPixels*3 {
				enemy.Damage(player.AttackPower())
				if enemy.Health() <= 0 {
					//deadEnemies[enemyIndex] = struct{}{}
					enemy.Kill()
					fmt.Println("Enemy killed")
				}
			}
		}

		if enemy.Collider.Overlaps(player.Collider) {
			if enemy.Attack() {
				player.Damage(enemy.AttackPower())
				fmt.Println("Player health: ", player.Health())
				if player.Health() <= 0 {
					fmt.Println("Player killed")
				}
			}
		}

		if enemy.IsDead() {
			delete(enemies, enemyIndex)
		}

		enemy.SetCurrentAnimation(
			int(enemy.GetChangeInX()),
			int(enemy.GetChangeInY()),
		)
		enemy.UpdateCurrentAnimation()
	}
}

func (c *Controller) PickupPotion(player *entities.Player, potions map[uint16]*entities.Potion) {
	for potionIndex, potion := range potions {
		if player.X >= potion.X && player.Y >= potion.Y {
			//fmt.Printf("Potion %d picked up\n", i)
			player.Inventory.StoreItem(potion)
			delete(potions, potionIndex)
		}
	}
}
