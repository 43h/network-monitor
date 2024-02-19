package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strconv"
)

const BuildVersion = "Time:2023-09-03\nVersion:0.0.2\nAuthor:cc"

// 主窗口大小
const MainWinMinWidth = 600
const MainWinMinHeight = 400

var searchTE *walk.LineEdit
var tableView *walk.TableView
var sbi *walk.StatusBarItem

type Item struct {
	index int
	ip    string
	delay string
}

type DataModel struct {
	walk.TableModelBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Item
}

func (m *DataModel) RowCount() int {
	return len(m.items)
}

func NewDataModel() *DataModel {
	m := new(DataModel)
	m.FlushRows()
	return m
}

func (m *DataModel) FlushRows() {
	// Create some random data.
	num := len(ips)
	if num == 0 {
		m.items = make([]*Item, 1)
		m.items[0] = &Item{index: 0}
	} else {
		m.items = make([]*Item, num)
		for i := range m.items {
			m.items[i] = &Item{
				index: i,
				ip:    ips[i].ipv4.To4().String(),
				delay: strconv.FormatInt(ips[i].avgRtt, 10) + " ms",
			}
		}
	}

	m.PublishRowsReset()
}

func (m *DataModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.index
	case 1:
		return item.ip
	case 2:
		return item.delay
	}

	panic("unexpected col")
}

func winmain() {
	newwin := new(walk.MainWindow)
	datamodel := NewDataModel()
	MainWindow{
		Title:    "网络状态检测",
		AssignTo: &newwin,
		MinSize:  Size{MainWinMinWidth, MainWinMinHeight},
		Layout:   VBox{},

		Children: []Widget{
			TableView{
				AssignTo:      &tableView,
				Model:         datamodel,
				StretchFactor: 2,
				Columns: []TableViewColumn{
					TableViewColumn{
						DataMember: "No.",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "IP",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "Delay",
						Alignment:  AlignCenter,
						Width:      128,
					},
				},
			},
		},
	}.Create()

	newwin.Run()
}
