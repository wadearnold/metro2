package segments

import (
	"gopkg.in/check.v1"
)

func (s *SegmentTest) TestTrailerRecord(c *check.C) {
	segment := NewTrailerRecord()
	_, err := segment.Parse(s.sampleTrailerRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.sampleTrailerRecord)
	c.Assert(segment.Description(), check.Equals, TrailerRecordDescription)
}

func (s *SegmentTest) TestTrailerRecordWithInvalidData(c *check.C) {
	segment := NewTrailerRecord()
	_, err := segment.Parse("ERROR" + s.sampleTrailerRecord)
	c.Assert(err, check.Not(check.IsNil))
}

func (s *SegmentTest) TestPackedTrailerRecord(c *check.C) {
	segment := NewPackedTrailerRecord()
	_, err := segment.Parse(s.samplePackedTrailerRecord)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.String(), check.Equals, s.samplePackedTrailerRecord)
	c.Assert(segment.Description(), check.Equals, PackedTrailerRecordDescription)
}

func (s *SegmentTest) TestPackedTrailerRecordWithInvalidData(c *check.C) {
	segment := NewPackedTrailerRecord()
	_, err := segment.Parse("ERROR" + s.samplePackedTrailerRecord)
	c.Assert(err, check.Not(check.IsNil))
}
