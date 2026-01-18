package resi

import "github.com/srv-api/web/dto"

func (s *resiService) TrackPackage(courier, awb string) (*dto.TrackingResponse, error) {
	return s.Repo.Track(courier, awb)
}
