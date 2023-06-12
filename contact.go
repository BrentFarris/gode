package ode

// #include <ode/ode.h>
import "C"

// Contact parameter types
const (
	Mu2CtParam      = C.dContactMu2
	AxisDepCtParam  = C.dContactAxisDep
	FDir1CtParam    = C.dContactFDir1
	BounceCtParam   = C.dContactBounce
	SoftERPCtParam  = C.dContactSoftERP
	SoftCFMCtParam  = C.dContactSoftCFM
	Motion1CtParam  = C.dContactMotion1
	Motion2CtParam  = C.dContactMotion2
	MotionNCtParam  = C.dContactMotionN
	Slip1CtParam    = C.dContactSlip1
	Slip2CtParam    = C.dContactSlip2
	RollingCtParam  = C.dContactRolling
	Approx0CtParam  = C.dContactApprox0
	Approx11CtParam = C.dContactApprox1_1
	Approx12CtParam = C.dContactApprox1_2
	Approx1NCtParam = C.dContactApprox1_N
	Approx1CtParam  = C.dContactApprox1
)

// SurfaceParameters represents the parameters of a contact surface.
type SurfaceParameters struct {
	// must always be defined
	Mode int
	Mu   float32

	// only defined if the corresponding flag is set in mode
	Mu2       float32
	Rho       float32
	Rho2      float32
	RhoN      float32
	Bounce    float32
	BounceVel float32
	SoftErp   float32
	SoftCfm   float32
	Motion1   float32
	Motion2   float32
	MotionN   float32
	Slip1     float32
	Slip2     float32
}

func (s *SurfaceParameters) fromC(c *C.dSurfaceParameters) {
	s.Mode = int(c.mode)
	s.Mu = float32(c.mu)
	s.Mu2 = float32(c.mu2)
	s.Rho = float32(c.rho)
	s.Rho2 = float32(c.rho2)
	s.RhoN = float32(c.rhoN)
	s.Bounce = float32(c.bounce)
	s.BounceVel = float32(c.bounce_vel)
	s.SoftErp = float32(c.soft_erp)
	s.SoftCfm = float32(c.soft_cfm)
	s.Motion1 = float32(c.motion1)
	s.Motion2 = float32(c.motion2)
	s.MotionN = float32(c.motionN)
	s.Slip1 = float32(c.slip1)
	s.Slip2 = float32(c.slip2)
}

func (s *SurfaceParameters) toC(c *C.dSurfaceParameters) {
	c.mode = C.int(s.Mode)
	c.mu = C.dReal(s.Mu)
	c.mu2 = C.dReal(s.Mu2)
	c.rho = C.dReal(s.Rho)
	c.rho2 = C.dReal(s.Rho2)
	c.rhoN = C.dReal(s.RhoN)
	c.bounce = C.dReal(s.Bounce)
	c.bounce_vel = C.dReal(s.BounceVel)
	c.soft_erp = C.dReal(s.SoftErp)
	c.soft_cfm = C.dReal(s.SoftCfm)
	c.motion1 = C.dReal(s.Motion1)
	c.motion2 = C.dReal(s.Motion2)
	c.motionN = C.dReal(s.MotionN)
	c.slip1 = C.dReal(s.Slip1)
	c.slip2 = C.dReal(s.Slip2)
}

// ContactGeom represents a contact point.
type ContactGeom struct {
	Pos    Vector3
	Normal Vector3
	Depth  float32
	G1     Geom
	G2     Geom
	Side1  int
	Side2  int
}

// NewContactGeom returns a new ContactGeom.
func NewContactGeom() *ContactGeom {
	return &ContactGeom{
		Pos:    NewVector3(),
		Normal: NewVector3(),
	}
}

func (g *ContactGeom) fromC(c *C.dContactGeom) {
	Vector(g.Pos).fromC(&c.pos[0])
	Vector(g.Normal).fromC(&c.normal[0])
	g.Depth = float32(c.depth)
	g.G1 = cToGeom(c.g1)
	g.G2 = cToGeom(c.g2)
	g.Side1 = int(c.side1)
	g.Side2 = int(c.side2)
}

func (g *ContactGeom) toC(c *C.dContactGeom) {
	Vector(g.Pos).toC((*C.dReal)(&c.pos[0]))
	Vector(g.Normal).toC((*C.dReal)(&c.normal[0]))
	c.depth = C.dReal(g.Depth)
	c.g1 = g.G1.c()
	c.g2 = g.G2.c()
	c.side1 = C.int(g.Side1)
	c.side2 = C.int(g.Side2)
}

// Contact represents a contact.
type Contact struct {
	Surface SurfaceParameters
	Geom    ContactGeom
	FDir1   Vector3
}

// NewContact returns a new Contact.
func NewContact() *Contact {
	return &Contact{
		FDir1: NewVector3(),
		Geom:  *NewContactGeom(),
	}
}

func (c *Contact) fromC(cc *C.dContact) {
	c.Surface.fromC(&cc.surface)
	Vector(c.FDir1).fromC(&cc.fdir1[0])
	c.Geom.fromC(&cc.geom)
}

func (c *Contact) toC(cc *C.dContact) {
	c.Surface.toC(&cc.surface)
	Vector(c.FDir1).toC(&cc.fdir1[0])
	c.Geom.toC(&cc.geom)
}
