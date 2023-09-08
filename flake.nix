{
  description = "cellular automata backend dev flake";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts"; 
  }; 

  outputs = inputs@{ self, nixpkgs, flake-parts, ... }: flake-parts.lib.mkFlake {inherit inputs;}{
  systems = [
        "x86_64-darwin"
        "x86_64-linux"
        "aarch64-darwin"
        "aarch64-linux"
  ];
  perSystem = {pkgs, system, self', inputs', config, ...}: { 
      
      devShells.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          gopls
          go_1_21
        ];  
      }; 
    };
  };
}
