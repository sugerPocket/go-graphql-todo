window.graphqlRequest = function(queryString) {
  return new Promise(function(resolve, reject) {
    fetch('/graphql?' + queryString)
      .then(function (result) {
        var reader = result.body.getReader();
        reader.read().then(function(result) {
          result = String.fromCharCode.apply(null, result.value);
          resolve(JSON.parse(result));
        });
      });
  });
}