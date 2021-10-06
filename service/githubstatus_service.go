package service

import (
	"githubstatus/constants"
	"strings"
)

func GetSystemSlugnameByClassAttribute(htmlClass string) string {
	switch true {
	case strings.Contains(htmlClass, "status-green"):
		return constants.OPERATIONAL_SLUGNAME
	case strings.Contains(htmlClass, "status-yellow"):
		return constants.DEGRADED_YELLOW_SLUGNAME
	case strings.Contains(htmlClass, "status-orange"):
		return constants.DEGRADED_ORANGE_SLUGNAME
	case strings.Contains(htmlClass, "status-red"):
		return constants.INCIDENT_SLUGNAME
	case strings.Contains(htmlClass, "status-blue"):
		return constants.MAINTENANCE_SLUGNAME
	}
	panic("Não foi encontrado nenhum slugname para esse attributo de classe 1633516592")
}

func GetSystemLabelBySlugname(slugname string) string {
	switch slugname {
	case constants.OPERATIONAL_SLUGNAME:
		return "Operational"
	case constants.DEGRADED_YELLOW_SLUGNAME:
	case constants.DEGRADED_ORANGE_SLUGNAME:
		return "Degraded"
	case constants.INCIDENT_SLUGNAME:
		return "Incident"
	case constants.MAINTENANCE_SLUGNAME:
		return "Maintenance"
	}

	panic("Não foi encontrado nenhum label para esse slugname 1633516869")
}
