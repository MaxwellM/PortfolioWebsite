describe('Sample test', function() {
    it('Condition is true', function() {
        expect('AngularJS').toBe('AngularJS');
    });
});

describe('angularJSExampleChartCtrl', function() {
    beforeEach(module('app'));

    var $controller;

    beforeEach(inject(function(_$controller_){
        $controller = _$controller_;
    }));

    describe('$scope.visitors', function() {
        it('Check the scope object', function() {
            var $scope = {};
            var controller = $controller('angularJSExampleChartCtrl', { $scope: $scope });
            expect($scope.visitors).toEqual([]);
        });
    });
});
