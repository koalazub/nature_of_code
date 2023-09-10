{
  description = "cellular automata backend dev flake";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts"; 
    templ.url = "github:a-h/templ";
  }; 

  outputs = inputs@{ self, nixpkgs, flake-parts, templ, ... }: flake-parts.lib.mkFlake {inherit inputs;}{
  systems = [
        "x86_64-darwin"
        "x86_64-linux"
        "aarch64-darwin"
        "aarch64-linux"
  ];
  imports = [
    inputs.flake-parts.flakeModules.easyOverlay
   ];
  perSystem = {pkgs, system, self', inputs', config, ...}: {  
      overlayAttrs = {
        inherit (config.packages) templ;
      };

      devShells.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          gopls
          go_1_21
          vscode-langservers-extracted 
        ];  
      }; 
    };
  };
}
