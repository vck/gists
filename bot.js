var aiml = require('aiml')
 
aiml.parse('yulan.aiml', function(err, topics){
  var engine = new aiml.AiEngine('Default', topics, {name: 'Jonny'});
  var responce = engine.reply({name: 'Billy'}, "ES KRIM", function(err, responce){
    console.log(responce);
  });
});