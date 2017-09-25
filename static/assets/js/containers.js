(function(win) {
  "use strict";

  var containers = {};
  win.TodoList['containers'] = containers;

  function isDateEqual(date_f, date_s) {
    return (
      date_f.getFullYear() === date_s.getFullYear() &&
      date_f.getMonth() === date_s.getMonth() &&
      date_f.getDate() === date_s.getDate()
    );
  }
  function createDate(date) {
    date = new Date(date);
    return new Date(date.getFullYear(), date.getMonth(), date.getDate());
  }
  var app = React.createClass({
    getInitialState: function() {
      return {
        list: [],
        editor: {},
        editorOpen: false
      };
    },
    componentWillMount: function() {
      var self = this;
      graphqlRequest('query={list{id,title,content,start,end}}')
        .then(function(result) {
          self.setState({ list: result.data.list });
        });
    },
    isInDate: function(date, todo) {
      var start = new Date(todo.start);
      var end = new Date(todo.end);
      var now = new Date(Date.now());
      if (start <= now && end > now) {
        todo.date = now;
        return isDateEqual(date, now);
      } else if (start < now) {
        todo.date = end;
        return isDateEqual(date, end);
      } else if (end >= now) {
        todo.date = start;
        return isDateEqual(date, start);
      }
    },
    partingTodoList() {
      var self = this;
      var list = this.state.list;
      var dividedList = [];
      var currentDate = null;
      var currentDividedList = [];
      list.forEach(function(todo) {
        if (currentDate === null) {
          currentDate = createDate(todo.start);
        }
        var isInDate = self.isInDate(currentDate, todo);
        if (isInDate) {
          currentDividedList.push(todo);
        } else {
          if (currentDividedList.length) {
            dividedList.push({
              date: currentDate,
              list: currentDividedList,
            });
          }
          currentDividedList = [];
          currentDate = new Date(todo.date);
          currentDividedList.push(todo);
        }
      });
      if (currentDividedList.length) {
        dividedList.push({
          date: currentDate,
          list: currentDividedList,
        });
      }
      return dividedList;
    },
    openEditor: function(todo) {
      this.setState({
        editor: {
          isOpen: true,
          todo: todo
        }
      });
    },
    render: function() {
      var dividedList = this.partingTodoList();
      var children = dividedList.map(function(child) {
        return React.createElement(TodoList.components.dailyTodoList, {
          date: child.date,
          list: child.list
        });
      });
      var isOpen = this.state.editor.isOpen;
      var todo = this.state.editor.todo;
      return React.createElement("div", null,
        React.createElement(TodoList.components.todoEditor, { isOpen: isOpen }),
        React.createElement(TodoList.components.topbar, { openEditor: this.openEditor || new Function() }),
        React.createElement.apply(React, ["div", null].concat(children))
      );
    }
  });
  containers.app = app;
})(window);
