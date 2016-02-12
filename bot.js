var aiml = require('aiml')
 
aiml.parseFile('yulan.aiml', function(err, topics){
  var engine = new aiml.AiEngine('Default', topics, {name: 'Jonny'});
  var responce = engine.reply({name: 'Billy'}, "Hi, dude", function(err, responce){
    console.log(responce);
  });
});