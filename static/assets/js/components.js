(function(win) {
  "use strict";
  
  // classes
  var components = {};
  win.TodoList['components'] = components;
  
  components.todoItem = React.createClass({
    propTypes: {
      id: PropTypes.string,
      start: PropTypes.string,
      end: PropTypes.string,
      title: PropTypes.string,
      content: PropTypes.string
    },
    defaultProps: {
      id: '',
      start: '',
      end: '',
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
        React.createElement('h3', { className: "title" },
          React.createElement('div', { className: "circle" }, null),
          React.createElement('span', null, content || ""),
          React.createElement('i', { className: "material-icons delete" }, 'delete'),
          React.createElement('i', { className: "material-icons edit" }, 'mode edit')
        ),
        React.createElement('span', { className: "content" }, content || "")
      );
    }
  });

  components.dailyTodoList = React.createClass({
    propTypes: {
      list: PropTypes.array
    },
    defaultProps: {
      date: new Date(Date.now()),
      list: []
    },
    render() {
      var list = this.props.list.map(function (todo) {
        return React.createElement(components.todoItem, todo);
      });
      var now = new Date(Date.now());
      now = new Date(now.getFullYear(), now.getMonth(), now.getDate());
      var date = new Date(this.props.date);
      date = new Date(date.getFullYear(), date.getMonth(), date.getDate());
      var expired = now > date;
      return React.createElement("div", { className: 'daily-list' + (expired ? ' expired' : '') },
        React.createElement("span", { className: 'day' }, moment(date).format("dddd DD[th] MMMM[,] YYYY")),
        React.createElement.apply(React, ["ul", { className: 'list' }].concat(list))
      )
    }
  });

  components.topbar = React.createClass({
    propTypes: {
      openEditor: PropTypes.func
    },
    render: function() {
      var openEditor = this.props.openEditor || new Function();
      return React.createElement("div", { className: "topbar clearfix" },
        React.createElement("h2", { className: "title float-left" }, "TODO",
          React.createElement("span", { onClick: openEditor, className: "add-a-todo float-right"}, "Add a todo +")
        )
      );
    }
  });
  
  components.todoEditor = React.createClass({
    // propTypes: {
    //   isOpen: propTypes.bool
    //   todo: PropTypes.object,
    //   onConfirm: PropTypes.function,
    // },
    componentWillReceiveProps: function(nextProps) {
      this.setState({
        isOpen: nextProps.isOpen,
        todo: nextProps.todo || {}
      });
    },
    handleChange: function(event) {
      var name = event.target.name;
      var todo = new Object(this.state.todo);
      todo[name] = event.target.value;
      this.setState({
        todo: todo
      });
    },
    handleTimeChange: function(name, time) {
      var time = moment(time);
      if (time.isValid()) {
        var todo = new Object(this.state.todo);
        todo[name] = time.format("YYYY/MM/DD HH:mm");
        this.setState({
          todo: todo
        });
      }
    },
    handleClose: function() {
      this.setState({
        isOpen: false,
        todo: {}
      });
    },
    handleSubmit: function() {
      if (this.isValid()) {
        this.props.onConfirm(this.state.todo);
        this.handleClose();
      };
    },
    isValid: function() {

    },
    render: function() {
      var state = this.state || {};
      var isOpen = state.isOpen || false;
      var todo = state.todo || {};
      var title = todo.title;
      var start = todo.start;
      var end = todo.end;
      var content = todo.content;
      return React.createElement(Reactstrap.Modal, { isOpen: isOpen, className: "todo-editor", toggle: this.handleClose },
        React.createElement(Reactstrap.ModalHeader, { toggle: this.handleClose }, "Edit a todo"),
        React.createElement(Reactstrap.ModalBody, null,
          React.createElement(Reactstrap.Form, null,
            React.createElement(Reactstrap.FormGroup, { className: "title", row: "" },
              React.createElement(Reactstrap.Input, { name: "title", placeholder: "Title", onChange: this.handleChange, value: title }, null)
            ),
            React.createElement(Reactstrap.FormGroup, { className: "time", row: "" },
              React.createElement(Datetime, { name: "start", className: "col-5 start", dateFormat: "YYYY/MM/DD", timeFormat: "HH:mm", onChange: this.handleTimeChange.bind(this, "start"), value: start, inputProps: { placeholder: 'Strat time' }}, null),
              React.createElement(Reactstrap.Label, { className: "col-2" }, "to"),
              React.createElement(Datetime, { name: "end", className: "col-5 end", dateFormat: "YYYY/MM/DD", timeFormat: "HH:mm", onChange: this.handleTimeChange.bind(this, "end"), value: end, inputProps: { placeholder: 'End time' } }, null)
            ),
            React.createElement(Reactstrap.FormGroup, { className: "content" },
              React.createElement(Reactstrap.Input, { name: "content", type: "textarea", placeholder: "Content", onChange: this.handleChange, value: content }, null)
            ),
          )
        ),
        React.createElement(Reactstrap.ModalFooter, null,
          React.createElement(Reactstrap.Button, { className: "submit", onClick: this.handleSubmit, color: "primary", block: true }, "OK")
        ),
      );
    }
  });
})(window);