import tkinter as tk
import random

# 踩地雷遊戲的主要邏輯
class MinesweeperGame:
    def __init__(self, rows, cols, num_mines):
        self.rows = rows
        self.cols = cols
        self.num_mines = num_mines
        self.board = [[0] * cols for _ in range(rows)]
        self.mines = set()
        self.revealed = set()
        self.game_over = False

    def generate_mines(self, first_click):
        possible_positions = [(x, y) for x in range(self.rows) for y in range(self.cols)]
        possible_positions.remove(first_click)
        self.mines = set(random.sample(possible_positions, self.num_mines))

    def calculate_adjacent_mines(self, row, col):
        if (row, col) in self.mines:
            return -1

        count = 0
        for dx in [-1, 0, 1]:
            for dy in [-1, 0, 1]:
                new_row = row + dx
                new_col = col + dy
                if 0 <= new_row < self.rows and 0 <= new_col < self.cols and (new_row, new_col) in self.mines:
                    count += 1

        return count

    def reveal_cell(self, row, col):
        if self.game_over:
            return

        if (row, col) in self.revealed:
            return

        self.revealed.add((row, col))

        if (row, col) in self.mines:
            self.game_over = True
            return

        if self.calculate_adjacent_mines(row, col) == 0:
            for dx in [-1, 0, 1]:
                for dy in [-1, 0, 1]:
                    new_row = row + dx
                    new_col = col + dy
                    if 0 <= new_row < self.rows and 0 <= new_col < self.cols:
                        self.reveal_cell(new_row, new_col)

    def is_game_won(self):
        return len(self.revealed) == (self.rows * self.cols) - self.num_mines

# 視窗化介面
class MinesweeperGUI:
    def __init__(self, rows, cols, num_mines):
        self.rows = rows
        self.cols = cols
        self.num_mines = num_mines
        self.game = MinesweeperGame(rows, cols, num_mines)

        self.window = tk.Tk()
        self.window.title("踩地雷遊戲")

        self.buttons = []
        for row in range(rows):
            button_row = []
            for col in range(cols):
                button = tk.Button(self.window, width=2, command=lambda r=row, c=col: self.click_cell(r, c))
                button.grid(row=row, column=col, padx=1, pady=1)
                button_row.append(button)
            self.buttons.append(button_row)

    def click_cell(self, row, col):
        if not self.game.game_over and (row, col) not in self.game.revealed:
            self.game.generate_mines((row, col))
            self.game.reveal_cell(row, col)
            self.update_board()

            if self.game.is_game_won():
                tk.messagebox.showinfo("遊戲結束", "恭喜你贏得了遊戲！")
                self.window.destroy()

            if self.game.game_over:
                tk.messagebox.showinfo("遊戲結束", "你踩到地雷了！")
                self.window.destroy()

    def update_board(self):
        for row in range(self.rows):
            for col in range(self.cols):
                if (row, col) in self.game.revealed:
                    if (row, col) in self.game.mines:
                        self.buttons[row][col].config(text="*", state="disabled")
                    else:
                        count = self.game.calculate_adjacent_mines(row, col)
                        self.buttons[row][col].config(text=count, state="disabled")

# 創建遊戲視窗
game_gui = MinesweeperGUI(rows=10, cols=10, num_mines=15)
game_gui.window.mainloop() 