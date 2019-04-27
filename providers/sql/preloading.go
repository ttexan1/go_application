package sql

import (
	"strings"
)

type preloading func(map[string]bool, string, string) map[string]bool

type preloadSet struct {
	Name string
	Func preloading
}

func preloadTable(preloaded map[string]bool, origin, val string, fn preloading) map[string]bool {
	splitted := strings.Split(val, ".")
	column := splitted[len(splitted)-1]
	if column == origin || column == origin+"s" {
		return preloaded
	}
	for _, val := range splitted[:len(splitted)-1] {
		if column == val {
			return preloaded
		}
	}
	preloaded[val] = true
	if fn != nil {
		preloaded = fn(preloaded, origin, val+".")
	}
	return preloaded
}

func preloadAnswer(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadBranchesEvent(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadCategory(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadClientsEvent(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadEmail(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"EntorsEmails", preloadEntorsEmail},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}

func preloadEntorsScheduledEventDate(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadEntorsEntryEvent(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadEntorsEmail(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadEntryEvent(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"Event", preloadEvent},
		preloadSet{"EntorsEntryEvents", preloadEntorsEntryEvent},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}

func preloadEvent(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"BranchesEvents", preloadBranchesEvent},
		preloadSet{"Category", preloadCategory},
		preloadSet{"ClientsEvents", preloadClientsEvent},
		preloadSet{"Emails", preloadEmail},
		preloadSet{"EntryEvent", preloadEntryEvent},
		preloadSet{"PostApplicationLink", preloadPostApplicationLink},
		preloadSet{"EventsQuestions", preloadEventsQuestion},
		preloadSet{"ScheduledEvent", preloadScheduledEvent},
		preloadSet{"TargetYears", preloadTargetYear},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}

func preloadEventsQuestion(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"Question", preloadQuestion},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}

func preloadGraduationYear(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadPostApplicationLink(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadQuestion(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadScheduledEvent(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"Event", preloadEvent},
		preloadSet{"ScheduledEventDates", preloadScheduledEventDate},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}

func preloadScheduledEventDate(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"ScheduledEvent", preloadScheduledEvent},
		preloadSet{"ScheduledEventPlace", preloadScheduledEventPlace},
		preloadSet{"EntorsScheduledEventDates", preloadEntorsScheduledEventDate},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}

func preloadScheduledEventPlace(preloaded map[string]bool, origin, namespace string) map[string]bool {
	return preloaded
}

func preloadTargetYear(preloaded map[string]bool, origin, namespace string) map[string]bool {
	for _, p := range []preloadSet{
		preloadSet{"GraduationYear", preloadGraduationYear},
	} {
		preloaded = preloadTable(preloaded, origin, namespace+p.Name, p.Func)
	}
	return preloaded
}
