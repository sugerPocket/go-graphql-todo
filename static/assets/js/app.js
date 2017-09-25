(function(win) {
  "use strict";

  // classes
  var TodoList = {};
  win.TodoList = TodoList;

  win.onload = function() {
    ReactDOM.render(
      React.createElement(TodoList.containers.app),
      document.getElementById("app")
    );
  }
  
})(window);
