package myopengl

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera_Movement int

const (
	FORWARD Camera_Movement = iota
	BACKWARD
	LEFT
	RIGHT

	yaw         float32 = -90
	pitch       float32 = 0
	speed       float32 = 2.5
	sensitivity float32 = 0.1
	zoom        float32 = 45
)

type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3
	WorldUp  mgl32.Vec3

	Yaw              float32
	Pitch            float32
	MovementSpeed    float32
	MouseSensitivity float32
	Zoom             float32
}

func NewCamera(position, up mgl32.Vec3, yaw, pitch float32) *Camera {
	camera := &Camera{}
	camera.Front = mgl32.Vec3{0, 0, -1}
	camera.MovementSpeed = speed
	camera.MouseSensitivity = sensitivity
	camera.Zoom = zoom
	camera.Position = position
	camera.WorldUp = up
	camera.Yaw = yaw
	camera.Pitch = pitch
	camera.updateCameraVectors()
	return camera
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}

func (c *Camera) ProcessKeyBoard(direction Camera_Movement, deltaTime float32) {
	velocity := c.MovementSpeed * deltaTime
	switch direction {
	case FORWARD:
		c.Position = c.Position.Add(c.Front.Mul(velocity))
	case BACKWARD:
		c.Position = c.Position.Sub(c.Front.Mul(velocity))
	case LEFT:
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	case RIGHT:
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	}
}

func (c *Camera) ProcessMouseMovement(xOffset, yOffset float64, constrainPitch bool) {
	xOffset *= float64(c.MouseSensitivity)
	yOffset *= float64(c.MouseSensitivity)

	c.Yaw += float32(xOffset)
	c.Pitch += float32(yOffset)

	if constrainPitch {
		if c.Pitch > 89 {
			c.Pitch = 89
		} else if c.Pitch < -89 {
			c.Pitch = -89
		}
	}
	c.updateCameraVectors()
}

func (c *Camera) ProcessMouseScroll(yOffset float64) {
	c.Zoom -= float32(yOffset)
	if c.Zoom < 1 {
		c.Zoom = 1
	} else if c.Zoom > 45 {
		c.Zoom = 45
	}
}

func (c *Camera) updateCameraVectors() {
	var front mgl32.Vec3
	front[0] = float32(math.Cos(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch))))
	front[1] = float32(math.Sin(float64(mgl32.DegToRad(c.Pitch))))
	front[2] = float32(math.Sin(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch))))

	c.Front = front.Normalize()
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()

}
