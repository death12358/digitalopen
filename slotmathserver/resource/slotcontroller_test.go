package resource

import (
	"reflect"
	"testing"

	"digitalopen/games"
	"digitalopen/sg001"
)

func TestNew(t *testing.T) {
	type args struct {
		rtp    map[string]string
		game   *games.Games
		defRTP string
	}
	tests := []struct {
		name string
		args args
		want *SlotController
	}{
		// TODO: Add test cases.
		{name: "test1",
			args: args{
				rtp:    sg001_rtps,
				game:   games.NewGames(sg001.New()),
				defRTP: sg001_rtp,
			},
			want: &SlotController{
				rtp:        sg001_rtps,
				game:       games.NewGames(sg001.New()),
				code:       "SG001",
				defaultRTP: "98"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rtp, tt.args.game, tt.args.defRTP); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJP(t *testing.T) {
	type args struct {
		rtp        map[string]string
		game       *games.Games
		defRTP     string
		jp_def     *[]string
		jp_sorts   []string
		jp_mapping map[string]string
	}
	tests := []struct {
		name string
		args args
		want *SlotController
	}{
		// TODO: Add test cases.
		{name: "test1",
			args: args{
				rtp:        sg001_rtps,
				game:       games.NewGames(sg001.New()),
				defRTP:     sg001_rtp,
				jp_def:     &sg001_jp_def,
				jp_sorts:   sg001_jppools_sorts,
				jp_mapping: sg001_jp_mapping},
			want: &SlotController{
				rtp:        sg001_rtps,
				game:       games.NewGames(sg001.New()),
				code:       "SG001",
				defaultRTP: "98",
				jp_def:     sg001_jp_def,
				jp_sorts:   sg001_jppools_sorts,
				jp_mapping: sg001_jp_mapping}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJP(tt.args.rtp, tt.args.game, tt.args.defRTP, tt.args.jp_def, tt.args.jp_sorts, tt.args.jp_mapping); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJP() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestSlotController_ExecuteSpin(t *testing.T) {
// 	bet_jn := json.Number("50")
// 	type fields struct {
// 		rtp        map[string]string
// 		game       *games.Games
// 		code       string
// 		defaultRTP string
// 		jp_def     []string
// 		jp_sorts   []string
// 		jp_mapping map[string]string
// 	}
// 	type args_spin struct {
// 		argv args.SpinArgs
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args_spin
// 		want    *games.Rounds
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test1",
// 			fields: fields{
// 				rtp:        sg001_rtps,
// 				game:       games.NewGames(sg001.New()),
// 				code:       "SG001",
// 				defaultRTP: "98",
// 				jp_def:     sg001_jp_def,
// 				jp_sorts:   sg001_jppools_sorts,
// 				jp_mapping: sg001_jp_mapping},
// 			args: args_spin{
// 				argv: args.SpinArgs{
// 					Id:       "1234567890",
// 					Brand:    "test1",
// 					Username: "tsUser",
// 					Pickem:   []string{"50"},
// 					Currency: "GCN",
// 					Bet:      bet_jn,
// 				}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sf := &SlotController{
// 				rtp:        tt.fields.rtp,
// 				game:       tt.fields.game,
// 				code:       tt.fields.code,
// 				defaultRTP: tt.fields.defaultRTP,
// 				jp_def:     tt.fields.jp_def,
// 				jp_sorts:   tt.fields.jp_sorts,
// 				jp_mapping: tt.fields.jp_mapping,
// 			}
// 			got, err := sf.ExecuteSpin(tt.args.argv)
// 			if err != nil {
// 				t.Errorf("SlotController.ExecuteSpin() error = %v, got %v", err, got)
// 				return
// 			}
// 			// if (err != nil) != tt.wantErr {
// 			// 	t.Errorf("SlotController.ExecuteSpin() error = %v, wantErr %v", err, tt.wantErr)
// 			// 	return
// 			// }
// 			// if !reflect.DeepEqual(got, tt.want) {
// 			// 	t.Errorf("SlotController.ExecuteSpin() = %v, want %v", got, tt.want)
// 			// }
// 		})
// 	}
// }
