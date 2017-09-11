(function() {
  "use strict";

  // classes

  var todoItem = React.createClass({
    propTypes: {
      id: React.propTypes.string,
      ddl: React.propTypes.date,
      title: React.propTypes.string,
      content: React.propTypes.string
    },
    defaultProps: {
      id: '',
      ddl: new Date(Date.now()),
      title: '',
      content: ''
    },

    // 设置初始状态
    getInitialState: function() {
      return {
        expanding: false
      };
    },
    render: function() {
      var title = this.props.title;
      var content = this.props.content;
      var expanding = this.state.expanding;
      return React.createElement('li', {
          className: expanding ? "todo-item expanding" : "todo-item"
        },
        React.createElement('h3', null, title || ""),
        React.createElement('span', null, content || "")
      );
    }
  });

  var dailyTodoList = React.createClass({
    propTypes: {
      date: React.propTypes.date,
      list: React.propTypes.array
    },
    defaultProps: {
      date: new Date(Date.now()),
      list: []
    },
    render() {
      var list = this.props.list.map(function (todo) {
        return React.createElement(todoItem, todo);
      });
      return React.createElement("div", { className: 'daily-list' },
        React.createElement("span", { className: 'day' }, this.props.date),
        React.createElement.apply(React, ["ul", { className: 'list' }].concat(list))
      )
    }
  });
})();
