package segments

import (
	"gopkg.in/check.v1"
)

func (s *SegmentTest) TestN1Segment(c *check.C) {
	segment := NewN1Segment()
	_, err := segment.Parse(s.sampleN1Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleN1Segment)
	c.Assert(segment.Description(), check.Equals, N1SegmentDescription)
}

func (s *SegmentTest) TestN1SegmentWithInvalidData(c *check.C) {
	segment := NewN1Segment()
	_, err := segment.Parse(s.sampleN1Segment[2:])
	c.Assert(err, check.Not(check.IsNil))
}
