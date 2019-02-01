var ngModule = angular.module('app');

ngModule.controller('angularJSExampleTableCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.addCharacter = addCharacter;
    $scope.showSection= showSection;
    $scope.pushToDB = pushToDB;

    $scope.showAddCharacter = false;
    $scope.showViewCharacters = false;


    function addCharacter(name, species, born, associates, gender, affiliation, masters) {
        let obj = {
            name: name,
            born: born ,
            associated: splitResults(associates),
            gender: gender,
            affiliation: splitResults(affiliation),
            masters: splitResults(masters)
        };


        console.log("Character: ", obj);

        $http.post("/addCharacterToDB", obj).then(function (res) {
            let results = res.data;
            console.log("RESULTS: ", results);
        }, function(error) {
            alert("ERROR ADDING CHARACTER TO DB: ", error);
        });

    }

    function showSection(section) {
        console.log("SECTION: ", section);
        if (section) {
            $scope.showAddCharacter = true;
            $scope.showViewCharacters = false;
        } else if (!section) {
            $scope.showViewCharacters = true;
            $scope.showAddCharacter = false;
        } else {
            console.log("Not sure what you selected there?");
        }
    }

    function splitResults(results) {
        let split;

        split = results.split(',');

        return split
    }

    function pushToDB (name, species, born, associates, gender, affiliation, masters) {
        $http.post("/addCharacterToDB", obj).then(function (res) {
            let results = res.data;
            console.log("RESULTS: ", results);
        }, function(error) {
           alert("ERROR ADDING CHARACTER TO DB: ", error);
        });
    }

}]);
