package components

type CombatCapable interface {
	Health() int
	AttackPower() int
	IsAttacking() bool
	Attack() bool
	UpdateCooldowns()
	Damage(amount int)
	Kill()
	IsDead() bool
}

type BasicCombat struct {
	health      int
	attackPower int
	isAttacking bool
	isDead      bool
}

func NewBasicCombat(health, attackPower int) *BasicCombat {
	return &BasicCombat{
		health:      health,
		attackPower: attackPower,
		isAttacking: false,
		isDead:      false,
	}
}

func (b *BasicCombat) Health() int {
	return b.health
}

func (b *BasicCombat) AttackPower() int {
	return b.attackPower
}

func (b *BasicCombat) Damage(amount int) {
	b.health -= amount
}

func (b *BasicCombat) IsAttacking() bool {
	return b.isAttacking
}

func (b *BasicCombat) Attack() bool {
	b.isAttacking = true
	return true
}

func (b *BasicCombat) Kill() {
	b.isDead = true
}

func (b *BasicCombat) IsDead() bool {
	return b.isDead
}

func (b *BasicCombat) UpdateCooldowns() {}

type EnemyCombat struct {
	*BasicCombat
	attackCooldown      int
	timeSinceLastAttack int
}

func (e *EnemyCombat) Attack() bool {
	if e.timeSinceLastAttack >= e.attackCooldown {
		e.isAttacking = true
		e.timeSinceLastAttack = 0
		return true
	}
	return false
}

func NewEnemyCombat(health, attackPower, attackCooldown int) *EnemyCombat {
	return &EnemyCombat{
		BasicCombat:         NewBasicCombat(health, attackPower),
		attackCooldown:      attackCooldown,
		timeSinceLastAttack: 0,
	}
}

func (e *EnemyCombat) UpdateCooldowns() {
	e.timeSinceLastAttack += 1
	e.isAttacking = false
}
