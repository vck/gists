AIMLInterpreter = require('AIMLInterpreter');
var aimlInterpreter = new AIMLInterpreter({name:'WireInterpreter', age:'42'});
aimlInterpreter.loadAIMLFilesIntoArray(['./yulan.aiml']);

var callback = function(answer, wildCardArray, input){
    console.log(answer + ' | ' + wildCardArray + ' | ' + input);
};

aimlInterpreter.findAnswerInLoadedAIMLFiles('What is your name?', callback);
aimlInterpreter.findAnswerInLoadedAIMLFiles('My name is Ben.', callback);
aimlInterpreter.findAnswerInLoadedAIMLFiles('What is my name?', callback);
